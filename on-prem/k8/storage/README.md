# How to Setup NFS (Network File System) on RHEL/CentOS/Fedora

## Benefits of NFS
    NFS allows local access to remote files.
    It uses standard client/server architecture for file sharing between all *nix based machines.
    With NFS it is not necessary that both machines run on the same OS.
    With the help of NFS we can configure centralized storage solutions.
    Users get their data irrespective of physical location.
    No manual refresh needed for new files.
    Newer version of NFS also supports acl, pseudo root mounts.
    Can be secured with Firewalls and Kerberos.

## Pre-Requisite
    NFS Server: nfsserver.example.com with IP-10.150.16.171
    NFS Client : nfsclient.example.com with IP-10.150.16.172

## Common installations on both Server and Client
    yum install nfs-utils nfs-utils-lib
    yum install portmap (not required with NFSv4)
    /etc/init.d/portmap start
    /etc/init.d/nfs start
    chkconfig --level 35 portmap on
    chkconfig --level 35 nfs on
        
## Setting up NFS Server
    DIRNAME="vol1"
    mkdir -p /mnt/disk/$DIRNAME 
    chcon -h system_u:object_r:bin_t:s0 /mnt/disk/$DIRNAME
    chcon -Rt svirt_sandbox_file_t /mnt/disk/$DIRNAME
    chmod 777 /mnt/disk/$DIRNAME
### Configure Export directory
    vi /etc/exports
        #Allow only client IP        
        #/mnt/disk/vol1 10.150.16.172(rw,sync,no_root_squash) 
        #Allow All    
        #/mnt/disk/vol1 *(rw,sync,no_root_squash) 
    
    service nfs start
    service nfs status
    showmount -e 
## Setting up NFS Client
    showmount -e 
    mkdir app_home
    mount -t nfs 10.150.16.171:/mnt/disk/vol1  app_home/
    mount | grep nfs #verify
    
    #To mount an NFS directory permanently on your system across the reboots, we need to make an entry in “/etc/fstab“
    vi /etc/fstab
        10.150.16.171:/mnt/disk/vol1 app_home/  nfs defaults 0 0

## Test NFS 
### NFS Server
    cat > /mnt/disk/vol1/server.txt
### NFS Client
    cat > /app_home/client.txt    
## Removing the NFS Mount
    umount app_home
    df -h -F nfs
    
## Important Commands
    Some more important commands for NFS.
    
    showmount -e : Shows the available shares on your local machine
    showmount -e <server-ip or hostname>: Lists the available shares at the remote server
    showmount -d : Lists all the sub directories
    exportfs -v : Displays a list of shares files and options on a server
    exportfs -a : Exports all shares listed in /etc/exports, or given name
    exportfs -u : Unexports all shares listed in /etc/exports, or given name
    exportfs -r : Refresh the server’s list after modifying /etc/exports
    
## NFS Options 
    Some other options we can use in “/etc/exports” file for file sharing is as follows.
    
    ro: With the help of this option we can provide read only access to the shared files i.e client will only be able to read.
    rw: This option allows the client server to both read and write access within the shared directory.
    sync: Sync confirms requests to the shared directory only once the changes have been committed.
    no_subtree_check: This option prevents the subtree checking. When a shared directory is the subdirectory of a larger file system, nfs performs scans of every directory above it, in order to verify its permissions and details. Disabling the subtree check may increase the reliability of NFS, but reduce security.
    no_root_squash: This phrase allows root to connect to the designated directory.    