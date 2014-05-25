VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  ["1", "2", "3"].each do |n|
    config.vm.define "zk#{n}" do |zk_config|
      zk_config.vm.box = "hashicorp/precise64"
      zk_config.vm.network "private_network", :ip => "192.168.12.1#{n}"
      zk_config.vm.provision :shell, :path => "vm-zookeeper.sh", :args => n
    end
  end

  config.vm.define "go" do |go_config|
    go_config.vm.box = "hashicorp/precise64"
    go_config.vm.network "private_network", :ip => "129.168.12.10"
    go_config.vm.provision :shell, :path => "vm-go.sh"
  end
end
