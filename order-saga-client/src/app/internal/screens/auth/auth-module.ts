import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';

import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';

import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { AuthRouting } from './auth-routing';


const PRIMENG_MODULES = [
    InputTextModule,
    ButtonModule,
];

@NgModule({
    imports: [
        AuthRouting,
        ReactiveFormsModule,
        ...PRIMENG_MODULES,
    ],
    declarations: [
        LoginComponent,
        RegisterComponent,
    ],
})
export class AuthModule {}