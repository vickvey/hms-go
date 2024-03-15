# Hotel Management System CLI App

- This is a command-line interface (CLI) application for managing hotel bookings. The application allows users to make, view, and cancel bookings for both AC and non-AC rooms. It provides a simple and efficient way to manage hotel reservations.

## Features

- Booking Management: Users can make bookings for AC or non-AC rooms, providing their personal details such as name, age, and Aadhar number.
- Room Availability: The system maintains the availability of rooms and ensures that bookings are made only for vacant rooms.
- Cancellation: Users can cancel their bookings if needed, freeing up the room for others.
- Status Checking: Users can check the status of their bookings, including details such as room type, amount paid, etc.
- View All Bookings: Administrators can view all bookings made in the hotel, including guest details and room information.

## Prerequisites

- Go programming language installed on your system.
- GNU Make (GNU Make is usually installed by default on most Linux distributions.)

## For Simple Use ( For normal users )
- Just download the [release](https://github.com/vickvey/hms-go/releases/tag/v1.0.0) for your particular operating system.
- Example for windows:
  - Download the `hms-go-windows-amd64.exe` file from the release section.
  - After downloading, open and run the executable file.
  - It will ask you that it could harm your computer, Ignore the warning as it is general security check for all outside applications.
  - Enjoy the app!

## Build And Run ( For devs )

To install and run the application, follow these steps:

1. Clone the repository to your local machine:

```bash
git clone https://github.com/vickvey/hms-go
```

2. Navigate to the project directory:

```bash
cd hms-go
```

3. Compile and run the application:

```bash
make
```
4. Go to bulid folder and run the required build for your operating system.
  Example: For linux run  :
  ```bash
  ➜  build git:(main) ✗ ./hms-go-linux-amd64 
Hotel Management System CLI App
Copyright@2024 vickvey
Press Enter to continue...
  ```

## Usage

Upon running the application, you will be presented with a menu containing various options:

- Make a Booking: Choose this option to make a new booking. You will be prompted to provide details such as room type, name, age, and Aadhar number.
- Show Booking Status: Check the status of a booking by entering the room number. This option displays the details of the booking if it exists.
- Cancel a Booking: Cancel a booking by entering the room number. This option removes the booking from the system and frees up the room.
- Show All Bookings: View a list of all bookings made in the hotel, including guest details and room information.
- Exit Menu: Choose this option to exit the menu and terminate the application.

## Contributors

vickvey copyright@2024
[github](https://github.com/vickvey)

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

Special thanks to the Google UUID library for generating unique booking IDs.

## Support

For any queries or issues, please open an issue on GitHub.

## Disclaimer

This project is for educational purposes only and should not be used in production environments without proper testing and validation.
