#include <stdio.h>
#include <string.h>

int main(int argc, char **argv)
{
	char key[] = "eZPz_LTrace"; 
	char pass[100];
	int n = strlen(key);

	printf("Your password: ");
	scanf("%s", pass);

	

	if (strcmp(pass, key) == 0 ) {
		puts("YESSSS GOT IT");
	} else {
		puts("NOOOOOOOO!");
	}
	return 0;
}
//compile : g++ -no-pie strcmp.cpp -o strcmp
