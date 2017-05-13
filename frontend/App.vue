<template>
	<div>
		<div class="page">
			<h1>About Naive Bayes Classification</h1>
			  <p>During lecture we discussed Bayes' theorem, which describes the probability of an event based on some previous knowledge called the prior probability. Baysian Naive Classification is an algorithm in machine learning that uses Bayes' theorem to classify text, or sub-objects made up of sub objects. When categorizing text the sub-objects are words. The probability that a block of text is in a category is based on what words it contains and how many of each kind.</p>
				<p>
					<em>For a more detailed explanation of Baysian Naive Classification see: https://en.wikipedia.org/wiki/Naive_Bayes_classifier.</em>
				</p>
				<p>The following classifier pulls Abstracts and Titles from the PubMed api. The abstract and title (article) are displayed to the user and the user assigns the article to a category. The algorithm then learns based on the categories given. When the articles are refreshed the user should see the probabilities that an article is in a category next to the name of a category for each article.</p>
				<p>Although the program is relatively fast and efficient it is not too good at actually categorizing articles. However, I only manually categorized a few dozen articles in my tests. I will continue to work on a better application that allows for longer term data collection (this site drops all the data every so often) and takes into account where the words are (title, abstract…).</p>
				<p>In machine learning, naive Bayes classifiers are a family of simple probabilistic classifiers based on applying Bayes' theorem with strong (naive) independence assumptions between the features.</p>
			<div height="100vh">
				<img src="/src/theorem.png" height="100px"/>
			</div>
		</div>
		<div class="page">
			<h1>STEP 1: Load Data From PubMed</h1>
			<p>
				This step request articles from PubMed using a pubmed query ({{query}}). The pubmed api simply searches a database of articles by date added. The goal of this project is to filter these using a machine learning algorithm based on Naive Bayes classification.
			</p>
			<button @click="fetch()" v-if="show_fetch">Fetch</button>
			<button @click="updateItems()" v-if="!show_fetch">Update Items</button>
			<div>
				<div v-for="item in items" class="article">
					<h1 @click="item.show_body = !item.show_body">{{item.data.title}}</h1>
					<p v-if="item.show_body">{{item.data.body}}</p>
					<p v-if="item.show_body">{{item.data.categories}}</p>
				</div>
			</div>
		</div>
		<div class="page">
			<h1>STEP 2: Classify Articles</h1>
			<article >
				<h2>{{item.data.title}}</h2>
				<p>{{item.data.body}}</p>
			</article>
			<div>
				<div v-for="(score, cat) in item.data.categories">
					<input type="checkbox" :value="cat" :id="cat" v-model="item.categories">
					<label :for="cat">{{cat}}</label>
				</div>
				<input type="text" @keyup.enter="categories.push(newCategory)" v-model="newCategory" placeholder="Category"/>
				<button @click="submitClasses(); next(); set_default_categories()">Next</button>
			</div>
		</div>
		<div>
			<a href="http://github.com/crhntr/math352">Source code (github)</a>
		</div>
	</div>
</template>
<script>
	export default {
	  data() {
	    return {
				query: "talimogene laherparepvec  [All Fields]",
	      threshold: 0.02,
	      newCategory: '',
				categories: [],
				items: [],
				show_fetch: true,
				current_item_id: 0,
				show_score: 0,
	      item: {
	        data: {},
					id: 0,
	      },
				query: "talimogene laherparepvec [All Fields]"
	    }
	  },
	  created() {},
	  methods: {
			submitClasses() {
				if (this.item.categories.length > 0) {
					this.$http.patch("/api/item/"+this.current_item_id+"/classes", {
						"classes": this.item.categories,
					})
				}
			},
			updateItems() {
				this.$http.get("/api/items").then(response => {
					this.items = []
					for (let item of response.data.data) {
						item.body = item.body || "(no body)"
						if (Object.keys(item.categories).length === 0) {
							this.show_score = false
							item.categories['relevant'] = 0
							item.categories['irrelevant'] = 0
						}
						this.items.push({
							show_body: false,
							data: item,
							categories: []
						})
					}
					this.current_item_id = 0;
					this.item = this.items[this.current_item_id]
				}, response => {
					console.log(response)
				})
			},
			fetch() {
				this.show_fetch = false
				this.$http.post("/act/item/fetch", this.query, {params:{"days": this.weeks_back}}).then(response => {
					this.show_fetch = false
					this.updateItems()
				}, response => {
					this.show_fetch = true
				})
			},
 	    next() {
				this.current_item_id++
				if (this.current_item_id >= this.items.length) {
					this.current_item_id = 0
					alert("you have seen every article and now are starting again")
				}
				console.log("next: " + this.current_item_id)
				this.item = this.items[this.current_item_id]
	    },
			set_default_categories() {
				console.log(this.item)
				for(let cl in this.item.data.categories) {
					console.log(cl, this.item.data[cl])
				}
			},
			refreshCategory(category) {

			}
	  }
	}
</script>
