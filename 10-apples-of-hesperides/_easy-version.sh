if [ "$#" -ne 3 ] ; then
    echo "usage: sh $0 <user> <ip> <port>"
    exit 1
fi

ssh-keygen -t rsa -f ~/.ssh/id_rsa -q -N ""

ssh-copy-id -p $3 $1@$2

ssh -p $3 $1@$2 "mkdir ~/bonus && git init ~/bonus"
ssh -p $3 $1@$2 