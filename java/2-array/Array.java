import java.util.ArrayList;

public class Array {
  public static void main(String[] args) {
    int[] array = new int[5];
    array[0] = 1;
    array[1] = 2;
    array[2] = 3;
    array[3] = 4;
    array[4] = 5;

    // or array creation with initialization
    int[] array2 = { 1, 2, 3, 4, 5 };

    System.out.println("Array with fixed size");
    for (int i = 0; i < 5; i++) {
      System.out.printf("Index: %d and Element %d\n", i, array[i]);
      System.out.printf("Index: %d and Element %d\n", i, array2[i]);
    }

    System.out.println("Array with dynamic size");
    var arrayList = new ArrayList<Integer>();
    arrayList.add(1);
    arrayList.add(2);
    arrayList.add(3);
    arrayList.add(4);
    arrayList.add(5);
    for (int i = 0; i < arrayList.size(); i++) {
      System.out.printf("Index: %d and Element %d\n", i, arrayList.get(i));
    }
  }
}
