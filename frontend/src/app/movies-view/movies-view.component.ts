import { Component, OnInit } from '@angular/core';
import {Router, ActivatedRoute } from '@angular/router';
import {HttpClient} from '@angular/common/http';
 
export class Movie {
  Id : Number ;
  Vote_average : Number;
  Title : string;
  Poster_path : string;
  Overview : string;
  Release_date : string
};
 
@Component({
  selector: 'app-movies-view',
  templateUrl: './movies-view.component.html',
  styleUrls: ['./movies-view.component.css']
})
export class MoviesViewComponent implements OnInit {
  moviesReceived: Array<Movie> = new Array<Movie>();
  movies: Array<Movie> = new Array<Movie>();


  constructor(private router : Router, private activatedRoute: ActivatedRoute, private httpClient: HttpClient) {
 
 
  }
 
  getMovies() {
 
    var config = {
      headers:
          {
              'Content-Type': 'application/json'
          }
    }  
 
    this.httpClient.get('http://localhost:3000/api/getMovies').subscribe(
 
      res => {
       
        this.moviesReceived= res['Movies']
        for( var i = 0; i < 20; i++) {
          this.movies.push(this.moviesReceived[i]);
        }
      }
    );
 
  }
 
  reserveSeat( movieId: Number){
 
   window.location.href = "/reserve?movie_id=" + movieId
  }
 
  ngOnInit() {
    this.getMovies()
  }
 
}