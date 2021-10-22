import java.util.Scanner;

public class login
{
private static String Username = "AccessLogin";
private static String Password = "GoldenTicket";

public static void main(String[] args)
{
    System.out.println("Please enter username");
    Scanner in = new Scanner(System.in);
    String UN = in.nextLine();

    System.out.println("Please enter password");
    Scanner inn = new Scanner(System.in);
    String PW = inn.nextLine();

    if (UN.equals(Username) && PW.equals(Password))
    {
        System.out.println("User has logged in successfully!");
        System.out.println("LAMECTF{"+Password+"}");
    }
    else {
        System.out.println("You have entered the wrong credentials, please try again.");
    }
}
 }
