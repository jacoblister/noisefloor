#include "ClientMock.hpp"

#include <thread>
#include <iostream>

void ClientMock::init() {
}

result<bool> ClientMock::start() {
    return true;
}

result<bool> ClientMock::stop() {
    return true;
}