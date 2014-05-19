VAGRANTFILE_API_VERSION = "2"
Vagrant.configure("2") do |config|
  config.vm.define "go" do |go_config|
    go_config.vm.box = "hashicorp/precise64"
    go_config.vm.network "private_network", :ip => "129.168.12.10"
    go_config.vm.provision :shell, :path => "go.sh"
  end

  ["1", "2", "3"].each do |n|
    config.vm.define "zk#{n}" do |zk1_config|
      zk1_config.vm.box = "hashicorp/precise64"
      zk1_config.vm.network "private_network", :ip => "192.168.12.1#{n}"
      zk1_config.vm.provision :shell, :path => "zookeeper.sh", :args => n
    end
  end
end
