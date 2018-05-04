# Show usage method
if [ "$#" -ne 3 ] ; then
    echo "usage: sh $0 <user> <ip> <port>"
    echo "\nEXAMPLE: \n - user: hercules\n - ip:   localhost\n - port: 1313"
    exit 1
fi
# Prompt to create key
while true; do
    read -p "Do you want to generate an RSA key? (y/n)  " yn
    case $yn in 
    [Yy]* )
            # Generate SSH Key | -t Type | -f Filename | -q Silence | -N New Password
            clear; echo "\n ------------ Generating RSA Key -----------"
            ssh-keygen -t rsa -f ~/.ssh/id_rsa -q -N "" && echo "\n -- RSA key created!"
            # Copy Public Key
            echo "\n ---------- Copying RSA Key to VM ----------\n    (Password is required only this one time)\n"
            ssh-copy-id -i ~/.ssh/id_rsa -p $3 $1@$2
            break;;
    [Nn]* ) 
            break;;
    *) 
            echo "Please enter yes or no.\n";;
    esac
done

# Connect to VM using ssh and RSA Key & Create Bonus Repo
echo "\n ----- Connecting & Creating Bonus Repo -----";
ssh -p $3 $1@$2 "mkdir ~/bonus && git init ~/bonus"
ssh -p $3 $1@$2