#include "Process.hpp"

class ClientRESTServer {
  public:
    ClientRESTServer(Process &process) : process(process) { }
    void init();
    void run();
  private:
    Process &process;
};
