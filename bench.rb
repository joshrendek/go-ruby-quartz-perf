require 'bcrypt'
require 'benchmark/ips'
require 'quartz'

pass = 'password'

client = Quartz::Client.new(file_path: './test.go')
hash = BCrypt::Password.create("password").freeze
Benchmark.ips do |x|
  x.time = 20
  x.warmup = 2
  x.report("ruby-bcrypt") {
    hash == pass
  }

  x.report("go-bcrypt") {
    client[:bcrypter].call('Check', 'Password' => pass, 'Hash' => hash)
  }

  x.report("ruby-add") {
    a = 1+2
    a == 3
  }

  x.report("go-add") {
    client[:bcrypter].call('Add', {})
  }

  x.compare!
end
