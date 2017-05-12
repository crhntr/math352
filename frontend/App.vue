<template>
	<div>
		<div class="page">
			<h1>Load Data From PubMed</h1>
			<button @click="fetch()">Fetch</button>
		</div>
		<div class="page">
			<h1>About Bayesian</h1>
			<div height="100vh">
				<img src="/src/theorem.png" height="100px"/>
			</div>
			<strong>Naive Bayes classifier</strong>
			<p>
				In machine learning, naive Bayes classifiers are a family of simple probabilistic classifiers based on applying Bayes' theorem with strong (naive) independence assumptions between the features.
			</p>
		</div>
		<div class="page">
			<h1>Iterate through Articles and Classify</h1>
			<article >
				<h2>{{item.data.title}}</h2>
				<p>{{item.data.abstract}}</p>
			</article>
			<div>
				<div v-for="category in categories">
					<input type="checkbox" :value="category" :id="category" v-model="item.categories">
					<label :for="category">{{category}}</label>
				</div>
				<input type="text" @keyup.enter="categories.push(newCategory)" v-model="newCategory" placeholder="Category"/>
				<button @click="submitClasses(); next()">Next</button>
			</div>
		</div>
		<div class="page">
			<h1>Show Articles Per Category</h1>
			<a href="/api/class/relevant">class relevant</a>
			<a href="/api/class/relevant/items">class relevant items</a>
		</div>
	</div>
</template>
<script>
	export default {
	  data() {
	    return {
	      threshold: 0.5,
	      newCategory: '',
	      categories: ['relevant', 'irrelevent'],
	      item: {
	        data: {},
					id: 0,
	        categories: [],
	      },
				query: "talimogene laherparepvec [All Fields]"
	    }
	  },
	  created() {},
	  methods: {
			submitClasses() {
				this.$http.patch("/api/item/"+this.item.id+"/classes", {
					"classes": this.item.categories,
				})
			},
			fetch() {
				this.$http.post("/act/item/fetch", this.query, {params:{"days": this.weeks_back}})
			},
 	    next() {
	      this.$http.get("/act/item/next").then(response => {
					this.item.data = response.data.data;
					this.item.id = response.data.id;
	      }, response => {
	        alert(JSON.stringify(response));
	      });
	    }
	  }
	}
</script>
