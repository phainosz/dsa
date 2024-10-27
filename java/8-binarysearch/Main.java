import java.util.ArrayList;
import java.util.List;

public class Main {

  record DummyTest(int expected, int needle) {
  }

  public static void main(String[] args) {
    var main = new Main();

    int[] array = { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 };
    var tests = main.prepareValues(array);

    for (var test : tests) {
      var result = main.binarySearch(array, test.needle);

      main.compareIndexes(test.expected, result);
      System.out.printf("searched finnished successful for needle %d with index %d\n", test.needle, result);
    }

  }

  private List<DummyTest> prepareValues(int[] array) {
    var tests = new ArrayList<DummyTest>();
    tests.add(new DummyTest(-1, 999));
    tests.add(new DummyTest(-1, 1099));
    tests.add(new DummyTest(-1, 98989));
    for (int i = 0; i < array.length; i++) {
      tests.add(new DummyTest(i, array[i]));
    }

    return tests;
  }

  private int binarySearch(int[] arr, int needle) {
    var left = 0;
    var right = arr.length - 1;

    while (left <= right) {

      var mid = left + (right - left) / 2;

      if (arr[mid] == needle) {
        return mid;
      }

      if (arr[mid] < needle) {
        left = mid + 1;
      } else {
        right = mid - 1;
      }
    }

    return -1;
  }

  private void compareIndexes(int expected, int got) {
    if (expected != got)
      throw new RuntimeException("linear search failed");
  }
}
