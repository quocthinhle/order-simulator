import { NgModule } from '@angular/core';
import { Route, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';

const AUTH_ROUTES: Route[] = [
    {
        path: 'register',
        component: RegisterComponent,
    },
    {
        path: 'login',
        component: LoginComponent,
    },
]

@NgModule({
    imports: [
        RouterModule.forChild(AUTH_ROUTES),
    ],
})
export class AuthRouting {

}