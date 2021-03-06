require 'nokogiri'
require 'pathname'
require 'json'

def underscore(camel_cased_word)
  camel_cased_word.to_s.gsub(/::/, '/').
    gsub(/([A-Z]+)([A-Z][a-z])/,'\1_\2').
    gsub(/([a-z\d])([A-Z])/,'\1_\2').
    tr("-", "_").
    downcase
end

args = ARGV.clone
ifile = args.shift
ofile = args.shift

if ifile.nil? && ofile.nil?
  puts "Usage: xml_to_json.rb <input_file> <output_file>"
  exit 0
end

unless File.exist? ifile
  puts "Input file #{ifile} does not exist"
  exit -1
end


if File.exist?(ofile) && File.directory?(ofile)
  puts "Output file #{ofile} already exists and is a directoy"
  exit -1
end

ofile_dir = Pathname.new(ofile).parent
unless ofile_dir.exist?
  puts "Expected #{ofile_dir} directory to exist"
  exit -1
end

unless ofile_dir.directory?
  puts "Expected #{ofile_dir} to be a directory"
  exit -1
end

doc = File.open(ifile) do |f| 
  Nokogiri::XML(f) { |config| config.huge }
end

nodes = doc.xpath("/*/row")

output = []
nodes.each do |node|
  node_data = {}
  node.keys.each do |k|
    if %w(id post_type_id accepted_answer_id score view_count owner_user_id last_editor_user_id answer_count comment_count favorite_count user_id class post_id count excerpt_post_id wiki_post_id reputation views up_votes down_votes account_id vote_type_id).include? k
      node_data[underscore(k)] = node[k].to_i
    else
      node_data[underscore(k)] = node[k]
    end
  end
  output << node_data
end

File.open(ofile, "w") do |f|
  f.write JSON.pretty_generate(output)
end

