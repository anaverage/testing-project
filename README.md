# Network Multithreaded Service for StatusPage

## Description

Our network service is designed to enhance customer interaction and optimize support operations. We provide tools and APIs for automating SMS, MMS, voice calls, and email systems. With a growing customer base spanning 160 countries, the load on our support team has increased. To streamline support processes and enhance efficiency, we've created a dedicated page to inform users about the current status of our systems.

## Key Features

- **Status Information**: Get real-time updates on the status of our services, including SMS, MMS, voice calls, and email.
- **Alerts and Incident History**: Receive alerts about potential issues and access a comprehensive incident history.
- **Integration with StatusPage**: Seamlessly integrated with our StatusPage website, allowing users to independently check system performance.
- **Automatic Data Collection**: Data on system status is automatically collected by various departments within the company.

## Technical Details

- **Programming Language**: Developed using the Go programming language.
- **Network Service**: Utilizes a network service for storing and managing system status data. This service accepts network requests and returns information about service status.
- **Automated Data Collection**: System status data is automatically collected by various departments within the company and stored in automated systems.
- **Git Flow**: Adheres to the Git Flow branching model for streamlined development and support workflows.

## Installation and Running

To demonstrate the service's functionality, you need to generate some data. To do this, follow these steps:

Clone the repository to your computer:

```bash
git clone https://github.com/beValed/statuspage-service.git
```

After that, start the simulator by running `main.go` located in the `simulator` directory.

Then, to start the service, navigate to the `cmd` directory and run `main.go`.

## Contributing

We welcome contributions to the project. If you have suggestions for improvements or bug fixes, please submit a Pull Request.

## License

This project is licensed under the terms of the MIT License.
