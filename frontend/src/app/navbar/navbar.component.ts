import { Component, OnInit } from '@angular/core';
import { Router, NavigationEnd } from '@angular/router';


@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  constructor(  public router: Router) { }
  navigateToAboutView(){
        this.router.navigate(['/aboutview']);

  }
  navigateToMoviesView(){
        this.router.navigate(['/moviesview']);


  }
  ngOnInit() {
  }

}
