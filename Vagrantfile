VAGRANTFILE_API_VERSION = "2"
VM_IP = "192.168.12.10"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "phusion-open-ubuntu-14.04-amd64"
  config.vm.box_url = "https://oss-binaries.phusionpassenger.com/vagrant/boxes/latest/ubuntu-14.04-amd64-vbox.box"
  config.vm.network "private_network", :ip => VM_IP
  config.vm.provision :shell, :path => "vm.sh", :args => VM_IP
end
