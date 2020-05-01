# Log File Archiving Script by CRON
  
1. Setup is performed by writing a list of the necessary files for archiving in the “.env” configuration file.  
  
    Example:  
        ROOT_COMPRESSED_LOG_DIRECTORY = "/path/to/log/arch/dir" - archive files save location  
        LOG_FILES = "/path/to/logfile/location = file_name1.log; /path/to/logfile/location = file_name2 .log" - path to file and name of the log file  
  
2. Files from the list will be renamed with the prefix of the date and time of renaming (YYYY_MM_DD_HHIISS_file_name.log); zip compressed; archive files will be moved to the specified directory, files with a prefix will be deleted and new log files with permissions 0666 will be created.  
  
  
# Install GO

```bash
apt install snapd  
```
Restart your  session  
```bash  
snap install go --classic  
```
```bash
go version  
```
  
  
# Deploy  

```bash
git clone ...
```   
```bash
echo "GOPATH=$HOME/.go-logrotate" >> ~/.bashrc
```
```bash  
cd ~/go-logrotate
```
```bash  
go get -d ./
```  
  
  
# Build & run  
```bash
go build  
```
```bash
./go-logrotate
```  
  
# Add CRON job  

```bash
crontab -e
```  
```bash 
# Example of job definition:  
# .---------------- minute (0 - 59)  
# |  .------------- hour (0 - 23)  
# |  |  .---------- day of month (1 - 31)  
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...  
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat  
# |  |  |  |  |  
# *  *  *  *  * user-name  command to be executed
```  
```bash   
crontab -l
```
!!!
