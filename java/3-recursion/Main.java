public class Main {

  public static void main(String[] args) {
    int number = 4;
    System.out.printf("Factorial of %d is %d\n", number, factorial(number));
    System.out.printf("Fibonacci of %d term is %d\n", number, fibonacci(number));
    printFibbonaci();

  }

  static int factorial(int n) {
    if (n <= 1)
      return 1;

    return n * factorial(n - 1);
  }

  static int fibonacci(int n) {
    if (n <= 1)
      return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
  }
  
  static void printFibbonaci() {
      System.out.println("Fibbonaci sequence 30 elements.");
    for (int i = 0; i < 30; i++) {
      System.out.printf("(%d)\t", fibonacci(i));
    }
    System.out.println();
  }

}
