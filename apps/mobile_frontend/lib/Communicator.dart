class Communicator {
  static final Communicator _singleton = new Communicator._internal();

  factory Communicator() {
    return _singleton;
  }

  Communicator._internal();
}
