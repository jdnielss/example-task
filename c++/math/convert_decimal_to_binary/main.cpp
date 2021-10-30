//
//  Created by :
//  Erwiyono Yusuf
//  Evri Bagas Saputro
//  Christian Folma Dio
//  Bayu Rahman Adinata 01085210005
//

#include <iostream>
#include <time.h>

void getConvertToBinary(int number);

using namespace std;

int main(void) {
    
    string programName = "Decimal to Binary Converter";
    string commandPress = "Press r to random a number: ";
    
    string inputCharStart;
    int randomNumber = NULL;
    
    cout << programName << endl;
    
    do {
        
        cout << commandPress;
        cin >> inputCharStart;
        
    } while (inputCharStart != "r");
    
    srand(time(NULL));
    randomNumber = (rand() % 256);
    
    cout << endl;
    
    cout << "Your Random Number" << endl;
    cout << "In decimal form: " << randomNumber << endl;
    getConvertToBinary(randomNumber);
    
}

void getConvertToBinary(int number) {
    string binary = "";
    
        while(number >= 1) {
            binary = to_string(number % 2) + binary;
            number = number / 2;
        }
            cout<<"in binary form: ";

    for(int i = 8; i > binary.length();) {
        binary = "0" + binary;
    }
    
    cout << binary;

    cout << endl;
}
