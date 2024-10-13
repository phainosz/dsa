import java.lang.System;

public class Bitwise {
  public static void main(String[] args) {
    // a = 5(00000101), b = 9(00001001)
    int a = 5, b = 9;

    System.out.printf("a value: %d\n", a);
    System.out.printf("a in binary: %s\n", Integer.toBinaryString(a));
    System.out.printf("a in binary left shifting: %s\n", Integer.toBinaryString(a << 1));
    System.out.printf("a value after left shifting: %d\n\n", a << 1);
    System.out.printf("b value: %d\n", b);
    System.out.printf("b in binary: %s\n", Integer.toBinaryString(b));
    System.out.printf("b in binary left shifting: %s\n", Integer.toBinaryString(b << 1));
    System.out.printf("b value after left shifting: %d\n", b << 1);
  }
}