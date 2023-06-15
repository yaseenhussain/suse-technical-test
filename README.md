Note: I have never worked on cypress, golang or terraform. I use selenium for UI automation and python and pytest for other automations.
I have not followed any standard practices as i am busy with our product release. I just got few hours to work on it, i completed whatever was asked in the assignment, but i could have added more checks and made code more standarized, but due to time constraint, i couldn't.



**First assignment:**
1. Make sure you have installed cypress and Rancher using docker(sudo docker run --privileged -d --restart=unless-stopped -p 80:80 -p 443:443 -
-name=rancher-server --net=rancher rancher/rancher)
2. Place this file "[first-assignment-rancher-login.cy.js](https://github.com/yaseenhussain/suse-technical-test/blob/main/First-assignment/first-assignment-rancher-login.cy.js)" in integration/ folder or any folder you like
3. Open this file and change username, password and baseURl(make sure these details are correct, else test will fail) and then run this file using cypress

**Second assignment**
1. Install golang
2. open the file "[Second_assignment_login_rancher_using_golang_v1.go](https://github.com/yaseenhussain/suse-technical-test/blob/main/Second-assignment/Second_assignment_login_rancher_using_golang_v1.go)" and change the username,password and base host(line no 17)
3. go run Second_assignment_login_rancher_using_golang_v1.go this will print out the response having the token to access rancher APIs -> This is working correctly
4. below are the ones which are working partially. I tried using standard go test framework with rancher client and rancher sdk
   1. I tried the using standard go test framework but due to some import issues, not able to run it, but still add it in the repo. Please do have a look. It is returning 400 response code
   2. I tried another version using rancher client, but due to some import issue couldn't run it. attaching it for reference

**Third assignment**
1. Install terraform
2. open the file "[main.tf](https://github.com/yaseenhussain/suse-technical-test/blob/main/Third-assignment/main.tf)" and change the credentials file path, project name and zone.
3. creating a  "n1-standard-1" vm in delhi region. 
4. Run below commands in order
   1. terraform init
   2. terraform plan -out=suseplan
   3. terraform apply suseplan (make sure you have enabled compute engine API access using GCP console)