export namespace modeles {
	
	export class Create {
	    Code: number;
	    Id: number;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Create(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Id = source["Id"];
	        this.Error = source["Error"];
	    }
	}
	export class EInfEmp {
	    Enterprise: string;
	    WorkStartDate: string;
	    JobTitle: string;
	
	    static createFrom(source: any = {}) {
	        return new EInfEmp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enterprise = source["Enterprise"];
	        this.WorkStartDate = source["WorkStartDate"];
	        this.JobTitle = source["JobTitle"];
	    }
	}
	export class EinfGroup {
	    Id: number;
	    Well: number;
	    Speciality: string;
	    GClass: number;
	    Semester: number;
	    Number: number;
	
	    static createFrom(source: any = {}) {
	        return new EinfGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Well = source["Well"];
	        this.Speciality = source["Speciality"];
	        this.GClass = source["GClass"];
	        this.Semester = source["Semester"];
	        this.Number = source["Number"];
	    }
	}
	export class GetId {
	    Code: number;
	    Id: number;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new GetId(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Id = source["Id"];
	        this.Error = source["Error"];
	    }
	}
	export class Inf_AllGroup {
	    Code: number;
	    Groups: EinfGroup[];
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Inf_AllGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Groups = this.convertValues(source["Groups"], EinfGroup);
	        this.Error = source["Error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Student {
	    id: number;
	    full_name: string;
	    speciality: string;
	    group_num: number;
	    semester: number;
	    well: number;
	    g_class: number;
	
	    static createFrom(source: any = {}) {
	        return new Student(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.full_name = source["full_name"];
	        this.speciality = source["speciality"];
	        this.group_num = source["group_num"];
	        this.semester = source["semester"];
	        this.well = source["well"];
	        this.g_class = source["g_class"];
	    }
	}
	export class Inf_AllStudents {
	    Code: number;
	    Students: Student[];
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Inf_AllStudents(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Students = this.convertValues(source["Students"], Student);
	        this.Error = source["Error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Inf_Employer {
	    Code: number;
	    Employer: EInfEmp;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Inf_Employer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Employer = this.convertValues(source["Employer"], EInfEmp);
	        this.Error = source["Error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Inf_Student {
	    Code: number;
	    Student: Student[];
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Inf_Student(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Student = this.convertValues(source["Student"], Student);
	        this.Error = source["Error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Inf_Subject {
	    Code: number;
	    Subject: string[];
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Inf_Subject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Subject = source["Subject"];
	        this.Error = source["Error"];
	    }
	}
	export class PdfDock {
	    Code: number;
	    File: number[];
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new PdfDock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.File = source["File"];
	        this.Error = source["Error"];
	    }
	}
	export class Remove {
	    Code: number;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Remove(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Error = source["Error"];
	    }
	}
	
	export class Update {
	    Code: number;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Update(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Error = source["Error"];
	    }
	}

}

