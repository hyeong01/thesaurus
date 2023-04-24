# thesaurus
### 0. What is this?
- A thesaurus implementation portfolio
- Backend: Go
- Frontend: HTML/CSS/JS (with chart.js)
- DB: PostgreSQL
- Publication: Docker

### 1. How to Use?
- Docker must me installed
- run sh files
  0. Depending on the version of your shell, choose between the winpty and the normal version. Git-bash needs winpty, the normal version gernerally works.
  1. Run ```build.sh``` in the ```backend``` folder to build the image
  2. Run ```run.sh``` in the ```backend``` folder to instantiate the image for the server
  3. Run ```dbSetup.sh``` to instantiate the image for DB
- Access the ```localhost:8080```
- Search "apple" to see the chart result. Other cases do not work due to lack of data.
