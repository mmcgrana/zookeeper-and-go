Vagrant.configure("2") do |config|
  config.vm.box = "hashicorp/precise64"
  config.vm.network :forwarded_port, host: 2181, guest: 2181
  config.vm.provision :shell, :path => "bootstrap.sh"  
end
