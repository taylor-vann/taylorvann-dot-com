// Brian Taylor Vann
// taylorvann dot com

*******
SELinux weidness!

SELinux has your back. Docker / Podman are a constant security threat.
In a production environment, "instances" of your application should be stateless.
And a lot of these security issues can be mitigated through well designed containers.
A development environment has a much more difficult time.

I'm a fan of Fedora so I have to shake hands and play well with SELinux.

// List file permisionss

ls -lZ


// Recursively allow all files and directories
// to be accessed by "virtual machines" (docker / podman / VMs)  

chcon -Rt svirt_sandbox_file_t /var/foo


The above command will resolve most "permission denied" errors on SELinux.

If you think you're having an SELinux issue?

// turn off SELinux
setenforce 1

Does it work now? Good, your computer is super hackable turn it back on.

However, some instances will still need further permissions


// Nginx specific permissions
chcon -Rt httpd_sys_content_t /path/to/nginx
chcon -Rt httpd_sys_content_t /path/to/nginx/conf.d

chcon -Rt httpd_sys_content_t /var/www
chcon -Rt httpd_sys_content_t /var/nginx


// PostgreSQL specific permissions
// socket allowances
chcon -Rt postgresql_var_run_t /var/run/postgresql

// allow directory to be treated as a database
chcon -Rt postgresql_db_t /path/to/potential/database/location/

