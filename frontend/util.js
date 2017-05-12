function get_cookie(name) {
	let all_cookies = document.cookie.split('; ');
	for (let i = 0; i < all_cookies.length; i++) {
		let [cname, value] = all_cookies[i].split('=', 2);
		if (cname === name) {
			console.log('found cookie: ' + value);
			return value;
		}
	}
	return false;
}

export default {
	get_cookie,
}
