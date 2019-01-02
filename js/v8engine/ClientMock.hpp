#include "Process.hpp"

class ClientMock {
  public:
    ClientMock(Process &process) : process(process) { }
    void init();
    result<bool> start();
    result<bool> stop();
  private:
    Process &process;
};
