export namespace define {
	
	export class Connection {
	    identity: string;
	    name: string;
	    address: string;
	    port: string;
	    user: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.identity = source["identity"];
	        this.name = source["name"];
	        this.address = source["address"];
	        this.port = source["port"];
	        this.user = source["user"];
	        this.password = source["password"];
	    }
	}

}

