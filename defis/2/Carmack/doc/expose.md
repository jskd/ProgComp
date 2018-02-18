
# Programmation et calcul générique sur GPU #

## Sommaire ##

1. Qu'est-ce que la "programmation GPU" ?
2. Petite histoire
3. Les différentes bibliothèques
4. OpenCL

## 1. Qu'est-ce que la "programmation GPU" ? ##

*General-purpose computing on graphics processing units (GPGPU)*

Usage du processeur graphique (*Graphical Processor Unit*, ou *GPU*)
pour effectuer des calculs typiquement effectués par
le processeur. Cela permet également d'effectuer des calculs
en parallèle, du fait de la structure du GPU et de la manière dont
les tratements sont effectués par le GPU.

TODO : il faudra entrer un peu plus en détails sur le fonctionnement
d'une carte graphique (processeur, mémoire)

## 2. Petite histoire ##

Motivations : Calculs complexe sur un volume important de données
\+ ou - complexes en parallelisant les tâches et en exploitant le processeur graphique (GPU).

Exemples: Multiplication matricielle et calcul sur des vecteurs en
*N* dimensions.

Quelques expérimentations:

 - 2001 : Multiplicaton matricielle exploitant le GPU
 - 2005: Décomposition LU (TODO kézako ?) sur GPU
 (plus rapide que sur CPU) - **reférence nécessaire**

OpenGL et DirectX fournissait de quoi pouvoir faire ces calculs
(assez rudimentaire).

Des bibliothèques spécifiques ont ensuite été développée pour
répondre à cette problématique (CUDA, DirectCompute, OpenCL).


## 3. Les différentes bibliothèques ##

- DirectCompute (Microsoft)
- Compute Unified Data Architecture ([CUDA][]) (Nvidia)
- [OpenACC][]
- Open Compute Library ([OpenCL][]) (Khronos)

## 4. Open Computing Library (OpenCL) ##

TODO : A remplir

## Liens ##

- Calcul sur GPU - https://www.wikiwand.com/en/General-purpose_computing_on_graphics_processing_units#/Implementations

---
[CUDA]: https://developer.nvidia.com/cuda-zone
[OpenACC]: https://www.openacc.org/
[OpenCL]: https://www.khronos.org/opencl/
