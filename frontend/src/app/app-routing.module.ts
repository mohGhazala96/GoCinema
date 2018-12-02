import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ReserveComponent } from './reserve/reserve.component';
import { MoviesViewComponent } from './movies-view/movies-view.component' 

const routes: Routes = [
  {
    path: 'reserve',
    component: ReserveComponent
  }, 

  {
    path: 'moviesview' ,
    component: MoviesViewComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
