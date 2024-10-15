public class Main {

  public static void main(String[] args) {
    var stack = new Stack<Integer>(1);

    stack.print();
    stack.push(2);
    stack.push(3);
    stack.print();

    stack.pop();
    stack.print();

    stack.pop();
    stack.print();

    stack.push(8);
    stack.push(10);
    stack.print();
    stack.peek();

  }

  static class Stack<T> {

    Node<T> top;
    int heigth;

    public Stack(T value) {
      var node = new Node<>(value);
      top = node;
      heigth = 1;
    }

    /*
     * insert value at the top
     */
    void push(T value) {
      assert heigth >= 0;
      var node = new Node<>(value);

      if (heigth != 0) {
        node.next = top;
      }
      top = node;
      heigth++;
    }

    /*
     * remove the first item at the top
     */
    Node<T> pop() {
      assert heigth >= 0;
      assert top != null;
      if (heigth == 0)
        return null;

      var temp = top;
      top = top.next;
      temp.next = null;
      heigth--;
      return temp;
    }

    /*
     * look at the first item at the top
     */
    void peek() {
      assert heigth >= 0;
      assert top != null;
      if (heigth <= 0)
        return;

      System.out.printf("Peeked value: %d\n", top.value);
    }

    void print() {
      System.out.println("=========== Printing Stack ===========");
      Node<?> temp = top;
      while (temp != null) {
        System.out.println(temp.value);
        temp = temp.next;
      }
      System.out.println("=========== End printing Stack ===========");
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
