function deepCopy(obj) {
	if (typeof obj === 'object' && obj != null && !Array.isArray(obj)) {
		const result = {};
		for (key in obj) {
			result[key] = deepCopy(obj[key]);
		}
		return result;
	}  if (Array.isArray(obj)) {    return obj.map((el) => {      return deepCopy(el);    });  }  return obj;}