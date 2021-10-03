/*
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 */

package nilairataratamhs;

import java.util.Scanner;
/**
 *
 * @author PC
 */
public class Main {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        String nama;

        double tugas, uts, uas, absensi, na;

        Scanner input = new Scanner(System.in); // memanggil class Scanner

        System.out.println("=========================================");
        System.out.println("Program perhitungan nilai akhir mahasiswa");
        System.out.print("Masukan nama :");
        nama = input.nextLine(); // Menyimpan nama dari inputan diatas
        System.out.print("Nilai Tugas :");
        tugas = input.nextInt(); // Menyimpan nama dari inputan diatas
        System.out.print("Nilai UTS :");
        uts = input.nextInt(); // Menyimpan nama dari inputan diatas
        System.out.print("Nilau UAS :");
        uas = input.nextInt(); // Menyimpan nama dari inputan diatas
        System.out.print("Nilai absensi :");
        absensi = input.nextInt(); // Menyimpan nama dari inputan diatas


        na = (0.15 * tugas + 0.20 * uts + 0.30 * uas + 0.35 * absensi);

        System.out.println(nama+" mendapatkan nilai akhir : "+na);
        System.out.println("=========================================");

        /* Contoh Output :
            run:
            =========================================
            Program perhitungan nilai akhir mahasiswa
            Masukan nama :Hacktoberfest
            Nilai Tugas :70
            Nilai UTS :90
            Nilau UAS :80
            Nilai absensi :78
            Hacktoberfest mendapatkan nilai akhir : 79.8
            =========================================
            BUILD SUCCESSFUL (total time: 23 seconds)
         */
    }

}
