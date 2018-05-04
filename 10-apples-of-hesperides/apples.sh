# Show usage method
if [ "$#" -ne 3 ] ; then
    echo "usage: sh $0 <user> <ip> <port>"
    echo "\nEXAMPLE: \n - user: hercules\n - ip:   localhost\n - port: 1313"
    exit 1
fi
# Generate SSH Key | -t Type | -f Filename | -q Silence | -N New Password
clear; echo "\n ------------ Generating RSA Key -----------"
ssh-keygen -t rsa -f ~/.ssh/id_rsa -q -N ""
echo "\n -- RSA key created!"
# Copy Public Key
echo "\n ---------- Copying RSA Key to VM ----------\n    (Password is required only this one time)\n"
ssh-copy-id -p $3 $1@$2
# Connect to VM using ssh and RSA Key & Create Bonus Repo
echo "\n ----- Connecting & Creating Bonus Repo -----";
ssh -p $3 $1@$2 "mkdir ~/bonus && git init ~/bonus"
ssh -p $3 $1@$2