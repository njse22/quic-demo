# -- mode: ruby --
# vi: set ft=ruby

Vagrant.configure("2") do |config|
    
    config.vm.define "goserver" do |server| 
      server.vm.box = "ubuntu/jammy64"
      server.vm.hostname = "server"
      server.vm.network "public_network", ip: "192.168.88.100", bridge: "enp0s31f6"
      server.vm.provision "file", source: "./provision/go/server/", destination: "$HOME/"
      server.vm.provision :shell do |shell| 
        shell.path = "./provision/go/installgo.sh"
      end
    end 

  (1..2).each do |i|
    config.vm.define "goclient-#{i}" do |client|
      client.vm.box = "ubuntu/jammy64"
      client.vm.hostname = "goclient-#{i}"
      client.vm.network "public_network", ip: "192.168.88.10#{i}", bridge: "enp0s31f6"
      client.vm.provider "virtualbox" do |vb|
        vb.customize ["modifyvm", :id, "--memory", "512", "--cpus", 1, "--name", "goclient-#{i}" ]
      end
      client.vm.provision "file", source: "./provision/go/client/", destination: "$HOME/"
      client.vm.provision :shell do |shell|
        shell.path = "./provision/go/installgo.sh"
      end
    end
  end 
end 
