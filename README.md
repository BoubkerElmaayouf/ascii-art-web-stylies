# Ascii-Art-Web

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, ascii-art.

## Description

This project allows users to convert text into ASCII art using a web-based graphical user interface. Users can choose from different banners and submit their text to be transformed. The server handles the requests and returns the ASCII art to be displayed on the web page.

### Supported Banners
- shadow
- standard
- thinkertoy

## HTTP Endpoints

### GET /
Sends an HTML response for the main page.

### POST /ascii-art
Sends data (text and a banner) to the Go server.

## Main Page Features

- **Text Input**: For user to enter the text to be converted.
- **Banner Selection**: Radio buttons, select objects, or other elements to switch between banners.
- **Submit Button**: Sends a POST request to `/ascii-art` and displays the result on the page.

## HTTP Status Codes

- **200 OK**: Everything went without errors.
- **404 Not Found**: Templates or banners not found.
- **400 Bad Request**: Incorrect requests.
- **500 Internal Server Error**: Unhandled errors.

## Authors

- Author 1: Elmaayouf Bopubker
- Author 2: Yassine Ouzddou
- Author 3: Mohammed fadil

## Usage

### How to Run

1. Clone the repository:
    ```sh
    git clone https://github.com/belmaayo/ascii-art-web-stylies.git
    ```

2. Navigate to the project directory:
    ```sh
    cd ascii-art-web-stylies
    ```

3. Build and run the Go server:
    ```sh
    go build
    ./ascii-art-web-stylies
    ```

4. Open your web browser and navigate to `http://localhost:8080`.

## Implementation Details

### Algorithm

1. **Text Input Handling**: The input text from the user is captured using a form.
2. **Banner Selection**: The user selects a banner, which is captured via radio buttons or a dropdown menu.
3. **POST Request**: The input text and selected banner are sent to the server via a POST request.
4. **ASCII Art Generation**: The server processes the input text using the specified banner to generate the ASCII art.
5. **Result Display**: The generated ASCII art is displayed on the web page. This can be done by:
   - Navigating to the `/ascii-art` route to display the result on a separate page.
   - Displaying the result on the home page by appending the result.

### Error Handling

- The server returns appropriate HTTP status codes based on the request outcome, ensuring robust error handling and user feedback.

### Dependencies

- Go programming language
- HTML/CSS for the front-end
- Any additional Go packages used for web server implementation and ASCII art generation
