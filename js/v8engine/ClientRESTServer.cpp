#include "ClientRESTServer.hpp"

#include <iostream>
#include <unistd.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <string.h>

void ClientRESTServer::init() {
}

void ClientRESTServer::run() {
    struct sockaddr_in server_addr, client_addr;
    int server_fd, client_fd, client_len, read_len, write_len;
    char client_message[2048] = {};

    server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd == -1) {
        std::cerr << "create socket fail" << std::endl;
        exit(0);
    }

    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = INADDR_ANY;
    server_addr.sin_port = htons(8080);

    if (bind(server_fd,(struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        std::cerr << "bind socket fail" << std::endl;
        exit(0);
    }

    listen(server_fd, 3);

    while (true) {
        client_len = sizeof(struct sockaddr_in);
        client_fd  = accept(server_fd, (struct sockaddr *)&client_addr, (socklen_t*)&client_len);

        read_len = read(client_fd, client_message, sizeof(client_message) - 1);
        if (read_len > 0) {
            char *header_endpoint = strstr(client_message, "/");
            char *header_proto    = strstr(client_message, "HTTP/1.1");
            char *client_body     = strstr(client_message, "\r\n\r\n");

            if (header_endpoint && header_proto && client_body) {
                std::string endpoint(header_endpoint, header_proto - 1);
                std::string request(client_body + 4, client_message + read_len);

                std::cout << "received: " << endpoint << std::endl << request << std::endl;

                std::string query_response = this->process.query(endpoint, request);

                std::string response = "HTTP/1.1 200 OK\r\n";
                response += "Content-Length: " + std::to_string(query_response.length()) + "\r\n";
                response += "Connection: close\r\n\r\n";
                response += query_response + "\r\n";
                write_len = write(client_fd, response.c_str(), response.length());

                std::cout << "end: " << response << std::endl;
            }
        }

        shutdown(client_fd, SHUT_RDWR);
        close(client_fd);
    }
}