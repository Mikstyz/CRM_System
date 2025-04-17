export namespace modeles {
	
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

}

