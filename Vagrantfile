Vagrant.configure("2") do |config|
  config.vm.base_mac = nil
  config.vm.synced_folder ".", "/vagrant", disabled: true

  config.vm.provider "virtualbox" do |vb|
    vb.gui = false
    vb.memory = "1024"
    vb.cpus = 1
    vb.linked_clone = true
  end

  N = 4
  (1..N).each do |machine_id|
    config.vm.define "docker-#{machine_id}" do |n|
      n.vm.hostname = "docker-#{machine_id}"
      n.vm.network "private_network", ip: "192.168.78.#{20+machine_id}"
      n.vm.box = "debian/buster64"

      if machine_id == N
        n.vm.provision :ansible do |ansible|
          ansible.limit = "all"
          ansible.playbook = "site.yaml"
          ansible.raw_arguments = [ "-D"]
          #ansible.raw_arguments = [ "-D", "--tags=backend"]
          #ansible.raw_arguments = [ "-D", "--tags=master"]

          ansible.host_vars = {
            "docker-1" => {"node_name"=>"docker-1"},
            "docker-2" => {"http_port" => 12345, "node_name"=>"docker-2"},
            "docker-3" => {"http_port" => 14444, "node_name"=>"docker-3"},
            "docker-4" => {"http_port" => 15555, "node_name"=>"docker-4"},

          }
          ansible.groups = {
            "master-servers" => ["docker-1"],
            "backend-servers" => ["docker-2", "docker-3", "docker-4"],
          }
          ansible.extra_vars = { ansible_python_interpreter:"/usr/bin/python3" }
        end
      end
    end
  end
end
