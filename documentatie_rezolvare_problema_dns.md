# Solve problem with DNS in linux for a certain site

On my virtual machine, I have two network interfaces and for some reason, the internet worked on my operating system, but on my virtual machine, the problem was that if I type ping 8.8.8.8 it works, but if I type ping goole.com, it doesn't work, so the DNS is not working corectly. I found a solutions which could work temporary. Let's solve the problem for a certain site.

Steps
-

1. From /home/ixia, type **sudo vi /etc/hosts**
2. You will see in the first lines some IPV4 adresses. Add one line with the IPV4 address fro the site that you want, the press the space key and type the name of the site. In my case, I typed: **156.140.12.140  bitbucket.it.keysight.com**
3. Press *esc* and *wq* to save changes and quit.
4. Press **ping <name_of_site_given_in_the_file>** to check if it works. I typed **ping bitbucket.it.keysight.com**