# Show usage method
if [ "$#" -lt 3 ] ; then
    echo "usage: sh $0 <user> <ip> <port> <flags>\n"
    echo "FLAGS:\n -k     generate key\n -r     create git repo\n"
    echo "EXAMPLE: \n - user: hercules\n - ip:   localhost\n - port: 1313\n"
    echo "WARNING: If invalid RSA Key provided, standard password will be used.\n"    
    exit 1
fi
# Generate and Copy SSH Key
if [ "$4" = "-k" ] || [ "$5" = "-k" ] ; then
    clear; echo "\n ------------ Generating RSA Key -----------"
    # Generate SSH Key | -t Type | -f Filename | -q Silence | -N New Password    
    ssh-keygen -t rsa -f ~/.ssh/hercules_key -q -N "" && echo "\n -- RSA key created!"
    echo "\n ---------- Copying RSA Key to VM ----------\n    (Password is required only this one time)\n"
    # Copy Public Key    
    ssh-copy-id -i ~/.ssh/hercules_key -p $3 $1@$2
fi
# Connect to VM using ssh and RSA Key & Create Bonus Git Repo
if [ "$4" = "-r" ] || [ "$5" = "-r" ] ; then
    echo "\n ----- Creating Bonus Repo -----";
    ssh -i ~/.ssh/hercules_key -p $3 $1@$2 "mkdir ~/bonus && git init ~/bonus"
fi
# Connect to VM using ssh and RSA Key
echo "\n ----- Connecting via SSH w/ RSA -----";
ssh -i ~/.ssh/hercules_key -p $3 $1@$2