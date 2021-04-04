#include <stdio.h>
#include <stdlib.h>
#include <windows.h>

int main() {
    HANDLE hComm;

    hComm = CreateFileA("\\\\.\\COM3", GENERIC_READ | GENERIC_WRITE, 0, NULL, OPEN_EXISTING, 0, NULL);

    if (hComm == INVALID_HANDLE_VALUE) {
        printf("An error occurred when we tried to open the serial port...");
        return -1;
    } else {
        printf("Opened the serial port successfully!");
    }

    CloseHandle(hComm);
    return 0;
}
