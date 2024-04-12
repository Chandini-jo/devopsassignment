This repository contains the code for deploying and testing a scalable and secure static web application on AWS.

Directory Structure

The repository is organized as follows:

/my-app
    /infrastructure
        ec2.yaml
        static-web-app.yaml
    /configuration
        playbook.yaml
    /tests
        test_web_server.py
    README.md

	•	The /infrastructure directory contains AWS CloudFormation templates for provisioning the required AWS resources.
	•	The /configuration directory contains an Ansible playbook for configuring the web server on the EC2 instance.
	•	The /tests directory contains a Python script for testing the web server configuration.

Usage

	1.	Set up AWS resources: Run the CloudFormation templates in the /infrastructure directory to set up the necessary AWS resources. Replace the placeholders in the templates with your actual values before running them.
	2.	Configure the web server: Run the Ansible playbook in the /configuration directory to configure the web server on the EC2 instance. Replace the placeholders in the playbook with your actual values before running it. Make sure to run the playbook from a machine that has network access to the target EC2 instance.
	3.	Run tests: Run the Python script in the /tests directory to test the web server configuration. Replace the placeholders in the script with your actual values before running it. Make sure to run the tests from a location that has access to the deployed web server.
