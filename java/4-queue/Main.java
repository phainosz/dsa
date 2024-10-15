public class Main {

  public static void main(String[] args) {
    enqueueExample();
    dequeueExample();
  }

  private static void dequeueExample() {
    System.out.println("===========dequeueExample===========");
    var queue = new Queue<Integer>(1);
    queue.enqueue(2);
    queue.enqueue(3);
    queue.enqueue(4);

    System.out.println("Value dequeued: " + queue.dequeue().value);

    queue.printQueue();
    System.out.println("================================");
  }

  private static void enqueueExample() {
    System.out.println("===========enqueueExample===========");

    var queue = new Queue<Integer>(1);

    queue.enqueue(2);

    queue.printQueue();
    System.out.println("================================");
  }

  static class Queue<T> {
    Node<T> first;
    Node<T> last;
    int length;

    public Queue(T value) {
      var node = new Node<T>(value);
      first = node;
      last = node;
      length++;
    }

    int getLength() {
      return length;
    }

    void enqueue(T value) {
      var node = new Node<T>(value);

      if (length == 0) {
        first = node;
      } else {
        last.next = node;
      }
      last = node;

      length++;
    }

    Node<T> dequeue() {
      if (length == 0)
        return null;

      var temp = first;
      if (length == 1) {
        first = null;
        last = null;
      } else {
        first = first.next;
        temp.next = null;
      }

      length--;
      return temp;
    }

    public void printQueue() {
      System.out.println("===========printQueue===========");
      var temp = first;
      while (temp != null) {
        System.out.print(temp.value + " ");
        temp = temp.next;
      }
      System.out.println();
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
