// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';
import {backend} from '../models';

export function AddActivityFromPomo(arg1:model.Activity):Promise<number>;

export function GetPomoWeekReport(arg1:string):Promise<Array<model.SessionDbRow>>;

export function GetPomos(arg1:string):Promise<Array<model.SessionDbRow>>;

export function Greet(arg1:string):Promise<string>;

export function StartPomo(arg1:model.Session):Promise<number>;

export function UpdatePomoSecondsLeft(arg1:backend.UpdateSessionSecondsLeft):Promise<boolean>;