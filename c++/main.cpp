// Program to random prime number with c++ basic.

#include <iostream>
#include <time.h>
#include <iomanip>

using namespace std;

bool checkIsPrime(int number);

int main() {

    string input, valueConstantToInput = "r";
    int randomNumber = 0, counter = 0, countPrime = 0;
    bool flag, inputValidation = true;

    srand(time(NULL));
    
    while (inputValidation == true) {
        cout << "Press r to generate random number : ";
        cin >> input;

        if (input != valueConstantToInput) {
            inputValidation = true;
        } else {
            inputValidation = false;
            cout << "Your random number(s): " << endl;
            for (int i = 1; i <= 100; i++) {
                randomNumber = (rand() % 100) + 1;
                flag = checkIsPrime(randomNumber);
                
                if (flag) {
                    countPrime++;
                    if (countPrime == 1 && counter == 0) {
                        cout << "none" << endl;
                    } else {
                        cout << randomNumber;
                    }
                    break;
                } else {
                    cout << randomNumber << setw(3);
                    counter++;
                }
            }
            if (counter > 0) {
                cout << endl;
                cout << counter << " non-prime numbers are generated" << endl;
            }
        };
    }
    
    return 0;
};

bool checkIsPrime(int number) {
    
    int counter = 0;
    bool isPrime;
    
    for (int i = 1; i <= number; i++) {
        if (number % i == 0) {
            counter++;
        }
    }
    
    if (counter == 2) {
        isPrime = true;
    } else {
        isPrime = false;
    }
    
    return isPrime;
}
