public class Main {

  public static void main(String[] args) {
    appendExample();
    prependExample();
    removeFirstExample();
    removeLastExample();

  }

  static void appendExample() {
    System.out.println("-----------------------");
    System.out.println("Insert example");
    var linkedList = new LinkedList<Integer>(2);
    linkedList.append(3);
    linkedList.append(4);

    linkedList.print();
  }

  static void prependExample() {
    System.out.println("-----------------------");
    System.out.println("Prepend example");
    var linkedList = new LinkedList<Integer>(2);
    linkedList.append(3);

    // 2 => 3 =>
    linkedList.print();

    linkedList.prepend(1);

    // 1 => 2 => 3 =>
    linkedList.print();
  }

  static void removeFirstExample() {
    System.out.println("-----------------------");
    System.out.println("Remove first example");
    var linkedList = new LinkedList<Integer>(2);
    linkedList.append(3);

    // 2 item - returns 2
    System.out.printf("want %d, got %d\n", 2, linkedList.removeFirst().value);
    // 1 item - returns 3
    System.out.printf("want %d, got %d\n", 3, linkedList.removeFirst().value);
    // 0 item - returns null
    System.out.println(linkedList.removeFirst());
  }

  static void removeLastExample() {
    System.out.println("-----------------------");
    System.out.println("Remove last example");
    var linkedList = new LinkedList<Integer>(2);
    linkedList.append(3);

    // 2 items - return 3
    System.out.printf("want %d, got %d\n", 3, linkedList.removeLast().value);
    // 1 item - return 2
    System.out.printf("want %d, got %d\n", 2, linkedList.removeLast().value);
    // 0 items - return null
    System.out.println(linkedList.removeLast());
  }

  static class LinkedList<T> {
    Node<T> head;
    Node<T> tail;
    int length;

    public LinkedList(T value) {
      var node = new Node<T>(value);
      head = node;
      tail = node;
      length++;
    }

    int getLength() {
      return length;
    }

    void print() {
      var temp = head;

      System.out.println("=========== Printing Linked List ===========");
      while (temp != null) {
        System.out.printf(" ( %d ) ", temp.value);
        temp = temp.next;
      }
      System.out.println();
    }

    // add item to the end
    void append(T value) {
      var node = new Node<T>(value);

      if (length == 0) {
        head = node;
      } else {
        tail.next = node;
      }
      tail = node;
      length++;
    }

    // add item to the start
    void prepend(T value) {
      var node = new Node<T>(value);

      if (length == 0) {
        head = node;
        tail = node;
      } else {
        node.next = head;
        head = node;
      }
      length++;
    }

    Node<T> removeFirst() {
      if (length == 0)
        return null;

      var temp = head;
      head = head.next;
      temp.next = null;

      length--;

      if (length == 0) {
        tail = null;
      }

      return temp;
    }

    Node<T> removeLast() {
      if (length == 0)
        return null;

      var temp = head;
      var previous = head;

      while (temp.next != null) {
        previous = temp;
        temp = temp.next;
      }

      tail = previous;
      tail.next = null;
      length--;

      // check if length is 0 after remove item
      if (length == 0) {
        head = null;
        tail = null;
      }
      return temp;
    }

  }

  static class Node<T> {
    T value;
    Node<T> next;

    public Node(T value) {
      this.value = value;
    }

  }
}