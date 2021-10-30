#include <iostream>
#include <cmath>
#include <iomanip>

using namespace std;

int main(void) {
    double number = 0;
    
    int i = 0;
    double result = 0;
    
    string wordingInput = "Input Number : ";
    string wordingNumber = "Number";
    string squareRoot = "Square Root";

    do
    {
        cout << wordingNumber << endl;
        cin >> number;

        result = sqrt(number);

        cout << squareRoot << endl;
        cout << result << endl;
        
        i++;
    } while (i != 3);
}
