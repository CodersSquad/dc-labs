#include <stdio.h>

__global__ void holaCUDA(float e) {
  printf("Hola, soy el hilo %d del bloque %d con valor pi->%f\n", threadIdx.x,blockIdx.x,e);
}

int main(int argc, char **argv){
  holaCUDA<<<8,4>>>(3.1416);
  cudaDeviceReset(); //Esta llamada reinicializa el device
  return 0;
}