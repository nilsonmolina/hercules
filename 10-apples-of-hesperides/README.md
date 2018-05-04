# Labour 10: Apples of Hesperides
**Goal**  
Learn about ssh and RSA to connect to your virtual machine without a password.

**Mandatory**  
Generate a way to access your virtual machine without a password. If you use a password, you fail. If you turn in something very wrong, you fail. Turn in the right thing. Make sure you can use it to connect to a virtual machine.
> Do not ever give your private key to anyone, it’s private for god sake!

## **Commands to Run**   
**Nothing here yet... Stay posted**
```bash
$ 
```

## **Enable `ssh` in Virtualbox**
For this project, we will need to ssh into our vm.  At this point, we will still require a password, but let's do things one step at a time.
1. With your VM turned off, go to:  
    `Settings -> Network -> Adapter 1 (default: attached to NAT) -> Advanced -> Port Forwarding`.
1. Add a new port forwarding rule with the following settings:  
    **Name:** *hercules-apples*  
    **Host Port:** *1313*  
    **Guest Port:** *22*  

    _Note:_ *Guest port must be 22, all other values are up to you. Leave IP's blank.*
1. Power on your VM and connect to it using ssh:
    ```bash
    # ssh -p <host port> <user>@<ip-address>
    $ ssh -p 1313 nmolina@localhost 
    ```
### Troubleshooting
**If you get an `ssh_exhange_identification: Connection closed by remote host`, what worked for me was to install openssh-server on the guest vm using the following command:**
```bash
$ sudo apt-get install openssh-server
```

## Create a New Sudo User
We probably do not want to ssh and rsa with the root user, so we can create a secondary user to do all of this with. **[user:hercules - pass:school42]**
1. Use the `adduser` command to add a new user to your system. You will then be prompted to set the password and user information.
    ```bash
    $ adduser <username>

    Set password prompts:
    Enter new UNIX password:
    Retype new UNIX password:
    passwd: password updated successfully

    User information prompts:
    Changing the user information for username
    Enter the new value, or press ENTER for the default
        Full Name []:
        Room Number []:
        Work Phone []:
        Home Phone []:
        Other []:
    Is the information correct? [Y/n]
    ```
1. Use the `usermod` command to add the user to the sudo group.
    ```bash
    $ usermod -aG sudo <username>
    ```

1. Test sudo access by switching to the new user account using `su`, and then running a command that requires sudo access:
    ```bash
    # switch to the new user
    $ su - <username>
    # test sudo privileges
    $ sudo ls -la /root
    ```

Reference:  
https://www.digitalocean.com/community/tutorials/how-to-add-and-delete-users-on-ubuntu-16-04

## Generate RSA Key
First things first, let us create our RSA key.

1. Create the RSA Key Pair
    ```bash
    $ ssh-keygen -t rsa

    Generating public/private rsa key pair.
    Enter file in which to save the key (/nfs/2017/n/nmolina/.ssh/id_rsa):
    Enter passphrase (empty for no passphrase):
    Enter same passphrase again:
    Your identification has been saved in /nfs/2017/n/nmolina/.ssh/id_rsa.
    Your public key has been saved in /nfs/2017/n/nmolina/.ssh/id_rsa.pub.
    The key fingerprint is:
    SHA256:3wYYQ0xEZmMmS8LPGO981IUlcDC2HjIR3PUW0WD2fOA nmolina@e1z3r3p2.42.us.org
    The keys randomart image is:
    +---[RSA 2048]----+
    |   ...=B/++oB+.  |
    |    oo.@o=.=.=.. |
    |     *+ =. .o E .|
    |    . ++.=..   . |
    |     o .S .      |
    |      o .. o     |
    |       .  . o    |
    |           .     |
    |                 |
    +----[SHA256]-----+
    ```

1. Copy the public key to the VM. (Might have to change network adapter, see troubleshooting below)
    ```bash
    ssh-copy-id hercules@10.113.100.139
    ```
1. And now you can ssh into the VM.
    ```bash
    ssh hercules@10.113.100.139
    ```

### Troubleshooting
**If you have an issue copying the public key to the VM, you need to change your network adapter from NAT to Bridged.**
1. Change the network adapter to Bridged:
`Settings -> Network -> Adapter 1 -> Attached to: -> Bridged`.
1. Get the IP Address of the machine:
    ```bash
    $ ip addr show

    1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1
        link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
        inet 127.0.0.1/8 scope host lo
            valid_lft forever preferred_lft forever
        inet6 ::1/128 scope host
            valid_lft forever preferred_lft forever
    2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
        link/ether 08:00:27:18:a9:0e brd ff:ff:ff:ff:ff:ff
        inet 10.113.100.139/16 brd 10.113.255.255 scope global enp0s3
            valid_lft forever preferred_lft forever
        inet6 fe80::a00:27ff:fe18:a90e/64 scope link
            valid_lft forever preferred_lft forever
    ```
1. Now you can try to connect using the following:
    ```bash
    $ ssh hercules@10.113.100.139
    ```

**If you have an issue with the key not being found, then you probably gave it a name and did not take the default.  Simply include it in your calls.**
```bash
$ ssh-copy-id -i ~/.ssh/hercules hercules@10.113.100.139
```