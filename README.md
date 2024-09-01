--- 
 
## Program Overview 
 
this is a simple Reverse Shell that helps to make persistence using registry and schtasks. But the main advantage is that this code bypasses all antiviruses. Let's GOO!

![image](https://github.com/user-attachments/assets/4edef35e-d3f9-4a71-8c6c-1a8a35c29c63)
![image](https://github.com/user-attachments/assets/edb8d5cf-8ec6-4ee8-aa0e-a35b51a82c27)

### Functionalities 
 
1. **Disable Real-Time Monitoring**: 
   - The function  d1()  executes a PowerShell command to disable Windows Defender's real-time monitoring. This is intended to allow the program to run without interference from security software. 
   - The command requires administrative privileges to execute successfully. 
 
2. **Add Entries to Windows Registry**: 
   - The function  d2()  adds entries to various registry keys to ensure the specified executable ( WMIprvSe.exe ) runs at startup or during specific events. 
   - It operates on the following registry keys: 
     -  HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run  
     -  HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce  
     -  HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServices  
     -  HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce  
   - The function runs these commands concurrently using Goroutines, allowing for non-blocking execution. 
 
3. **Create Scheduled Tasks**: 
   - The function  d3()  creates a scheduled task that executes the specified executable ( WMIprvSe.exe ) every minute. 
   - This ensures that the executable is run regularly, providing persistence even if the system is restarted. 
 
4. **Check Configuration**: 
   - The function  e1()  checks if the IP address ( a1 ) and port ( a2 ) are correctly defined. If either is not configured, the program will exit with an error message. 
 
5. **Establish TCP Connection**: 
   - The function  e2()  attempts to establish a TCP connection to the specified IP address and port. It will keep trying every 5 seconds until a connection is successful. 
   - Once connected, it returns the connection object. 
 
6. **Execute Commands via TCP**: 
   - The function  e3(c net.Conn)  starts a command shell ( cmd.exe ) and redirects its input, output, and error streams to the established TCP connection. This allows remote command execution through the established connection. 
 
7. **Main Functionality**: 
   - The  main()  function orchestrates the execution of all the above functionalities in sequence: 
     - It disables real-time monitoring. 
     - Adds registry entries for persistence. 
     - Creates scheduled tasks. 
     - Checks the configuration for the connection. 
     - Establishes a TCP connection. 
     - Starts a command shell that communicates over the TCP connection. 
 
### Important Notes 
 
- **Administrative Privileges**: The program requires administrative rights to modify the registry and create scheduled tasks. 
- **Persistence Mechanism**: The use of registry entries and scheduled tasks allows the program to maintain persistence across reboots. 
- **Potential Security Risks**: This program may be used for malicious purposes. Ensure that it is used responsibly and in compliance with applicable laws and regulations. 
 
### Dependencies 
 
- Go programming language (version 1.16 or higher recommended). 
 
### Usage 
 
1. Clone the repository. 
2. Ensure you have Go installed on your machine. 
 
--- 
