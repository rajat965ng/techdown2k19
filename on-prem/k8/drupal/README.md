wget https://www.apachefriends.org/xampp-files/7.4.2/xampp-linux-x64-7.4.2-0-installer.run
chmod +x xampp-linux-x64-1.8.3-3-installer.run
./xampp-linux-x64-1.8.3-3-installer.run

    ```
    Welcome to the XAMPP Setup Wizard.
    
    ----------------------------------------------------------------------------
    Select the components you want to install; clear the components you do not want
    to install. Click Next when you are ready to continue.
    
    XAMPP Core Files : Y (Cannot be edited)
    
    XAMPP Developer Files [Y/n] :y
    
    Is the selection above correct? [Y/n]: y
    
    ----------------------------------------------------------------------------
    Installation Directory
    
    XAMPP will be installed to /opt/lampp
    Press [Enter] to continue :
    
    ----------------------------------------------------------------------------
    Setup is now ready to begin installing XAMPP on your computer.
    
    Do you want to continue? [Y/n]: y
    
    ----------------------------------------------------------------------------
    Please wait while Setup installs XAMPP on your computer.
    
    Installing
    0% ______________ 50% ______________ 100%
    #########################################
    
    ----------------------------------------------------------------------------
    Setup has finished installing XAMPP on your computer.
    ```

    # vim /opt/lampp/etc/extra/httpd-xampp.conf
    
    <LocationMatch "^/(?i:(?:xampp|security|licenses|phpmyadmin|webalizer|server-status|server-info))">
    # Require local
    Require all granted
    ErrorDocument 403 /error/XAMPP_FORBIDDEN.html.var
    </LocationMatch>
    
    
    
    
sudo /opt/lampp/lampp restart    


mkdir /opt/lampp/htdocs/drupal
wget https://www.drupal.org/download-latest/tar.gz
tar -zxvf tar.gz
mv drupal-8.8.2/* /opt/lampp/htdocs/drupal/
sudo cp /opt/lampp/htdocs/drupal/sites/default/default.settings.php /opt/lampp/htdocs/drupal/sites/default/settings.php

cd /opt/lampp/htdocs
ps -ef | grep lampp [check the name of the group like daemon]
sudo chgrp daemon drupal -R
sudo chmod 775 drupal -R
sudo chown -R daemon:daemon drupal/    


/opt/lampp/etc/httpd.conf

cd /opt/lampp/htdocs/drupal && vim .htaccess
   
and Type following code:
    
    <IfModule mod_rewrite.c>
    RewriteEngine On
    RewriteCond %{REQUEST_FILENAME} !-f
    RewriteCond %{REQUEST_FILENAME} !-d
    RewriteRule ^(.*)$ index.php/$1 [L]
    </IfModule>
    
sudo /opt/lampp/lampp restart    
    