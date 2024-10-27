import java.util.ArrayList;
import java.util.List;

public class Main {

  record DummyTest(int expected, int needle) {
  }

  public static void main(String[] args) {

    int[] array = new int[] { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 };
    var tests = prepareValues(array);

    for (var test : tests) {
      var result = linearSearch(array, test.needle);

      compareIndexes(test.expected, result);
      System.out.printf("searched finnished successful for needle %d with index %d\n", test.needle, result);
    }

  }

  private static List<DummyTest> prepareValues(int[] array) {
    var tests = new ArrayList<DummyTest>();
    tests.add(new DummyTest(-1, 999));
    tests.add(new DummyTest(-1, 1099));
    tests.add(new DummyTest(-1, 98989));
    for (int i = 0; i < array.length; i++) {
      tests.add(new DummyTest(array[i], i));
    }

    return tests;
  }

  private static int linearSearch(int[] array, int needle) {
    for (int i = 0; i < array.length; i++) {
      if (array[i] == needle) {
        return i;
      }
    }
    return -1;
  }

  private static void compareIndexes(int expected, int got) {
    if (expected != got)
      throw new RuntimeException("linear search failed");
  }
}